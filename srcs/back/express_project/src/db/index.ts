// src/db/index.ts

import { Pool, QueryResult, QueryResultRow } from 'pg';

interface DatabaseConfig {
    user: string;
    host: string;
    database: string;
    password: string;
    port: number;
}

// データベース操作用のクラス
export class Database {
    private pool: Pool;

    constructor() {
        const config = this.validateConfig();
        this.pool = new Pool({
            ...config,
            max: 20,
            idleTimeoutMillis: 30000,
            connectionTimeoutMillis: 2000,
        });
        // エラーイベントのハンドリング
        this.pool.on('error', (err) => {
            console.error('Unexpected error on idle client', err);
        });
    }

    private validateConfig(): DatabaseConfig {
        const requiredEnvVars = [
            'POSTGRES_USER',
            'POSTGRES_HOST',
            'POSTGRES_DB',
            'POSTGRES_PASSWORD'
        ];
        const missingEnvVars = requiredEnvVars.filter(varName => !process.env[varName]);
        if (missingEnvVars.length > 0) {
            throw new Error(`Missing required environment variables: ${missingEnvVars.join(', ')}`);
        }

        return {
            user: process.env.POSTGRES_USER!,
            host: process.env.POSTGRES_HOST!,
            database: process.env.POSTGRES_DB!,
            password: process.env.POSTGRES_PASSWORD!,
            port: 5432
        };
    }

    async testConnection(): Promise<void> {
        try {
            const client = await this.pool.connect();
            await client.query('SELECT NOW()');
            client.release();
        } catch (error) {
            throw new Error(`Database connection test failed: ${error}`);
        }
    }

    async query<T extends QueryResultRow>(
        text: string,
        params?: any[]
    ): Promise<QueryResult<T>> {
        const client = await this.pool.connect();
        try {
            return await client.query(text, params);
        } catch (error) {
            throw new Error(`Query execution failed: ${error}`);
        } finally {
            client.release();
        }
    }

    async executeTransaction<T>(
        callback: (client: any) => Promise<T>
    ): Promise<T> {
        const client = await this.pool.connect();
        try {
            await client.query('BEGIN');
            const result = await callback(client);
            await client.query('COMMIT');
            return result;
        } catch (error) {
            await client.query('ROLLBACK');
            throw error;
        } finally {
            client.release();
        }
    }

    async end(): Promise<void> {
        await this.pool.end();
    }
}

export default new Database();