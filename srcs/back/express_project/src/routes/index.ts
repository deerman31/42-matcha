import { Router } from "express";
import { Database } from "../db";

export function setupRoutes(db: Database) {
    const router = Router();

    router.get("/", (req, res) => {
        res.send("Hello World!");
    });
    return router;
}