// app/routes/my-profile.tsx
import { useLoaderData, useNavigation, useNavigate, useFetcher } from "@remix-run/react";
import { redirect } from "@remix-run/node";
import type { LoaderFunction, ActionFunction } from "@remix-run/node";
import { getAccessToken, destroySession } from "~/session.server";
import { useEffect, useState } from "react";

interface UserInfo {
    username: string;
    email: string;
    lastname: string;
    firstname: string;
    birthdate: string;
    gender: 'male' | 'female';
    sexuality: 'male' | 'female' | 'male/female';
    area: string;
    self_intro: string;
    tags: string[];
    is_gps: boolean;
    latitude: number;
    longitude: number;
    fame_rating: number;
}

interface LoaderData {
    my_info: UserInfo;
    token?: string; // トークンをローダーデータに追加
    error?: string;
}

// セッション破棄用のアクション関数を追加
export const action: ActionFunction = async ({ request }) => {
    await destroySession(request);
    return redirect("/login");
};

export const loader: LoaderFunction = async ({ request }) => {
    // セッションからトークンを取得
    const token = await getAccessToken(request);
    if (!token) {
        return redirect("/login");
    }

    try {
        // ユーザー情報の取得
        const userResponse = await fetch("http://back:3000/api/my-profile", {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });

        if (!userResponse.ok) {
            if (userResponse.status === 401) {
                await destroySession(request);
                return redirect("/login");
            }
            throw new Error('Failed to fetch user data');
        }

        const userData = await userResponse.json();
        //return userData.my_info; // my_info を含むオブジェクトがそのまま返される
        return {
            ...userData,
            token // トークンをレスポンスに含める
        };

    } catch (error) {
        console.error('Error fetching profile data:', error);
        throw new Error('Failed to load profile data');
    }
};

export default function MyProfilePage() {
    const data = useLoaderData<LoaderData>();
    const [profileImage, setProfileImage] = useState<string>('');
    const navigation = useNavigation();
    const isLoading = navigation.state === "loading";
    //const navigate = useNavigate();  // useNavigateを追加
    const fetcher = useFetcher();  // セッション破棄用のfetcherを追加

    // document へのアクセスは useEffect 内で行う
    useEffect(() => {
        const fetchProfileImage = async () => {
            if (!data || !data.my_info || !data.token) {
                return;
            }

            try {
                const response = await fetch("/api/users/get/image1", {
                    headers: {
                        Authorization: `Bearer ${data.token}`,
                    },
                });

                if (!response.ok) {
                    if (response.status === 401) {
                        // セッション破棄のPOSTリクエストを送信
                        fetcher.submit(null, { method: "post" });
                        //navigate("/login");  // redirect の代わりに navigate を使用
                        return;
                    }
                    throw new Error('Failed to fetch user data');
                }
                const imageData = await response.json();
                setProfileImage(imageData.image);
            } catch (error) {
                console.error('Error fetching profile image:', error);
            }
        };

        fetchProfileImage();
    }, [data, fetcher]); // 初回レンダリング時のみ実行



    // データチェックを追加
    if (!data || !data.my_info) {
        return (
            <div className="min-h-screen bg-gray-100 flex items-center justify-center">
                <div className="text-xl font-semibold text-red-600">
                    プロフィールデータの読み込みに失敗しました
                </div>
            </div>
        );
    }
    // エラーチェック
    if (data.error) {
        return (
            <div className="min-h-screen bg-gray-100 flex items-center justify-center">
                <div className="text-xl font-semibold text-red-600">{data.error}</div>
            </div>
        );
    }


    if (isLoading) {
        return (
            <div className="min-h-screen bg-gray-100 flex items-center justify-center">
                <div className="text-xl font-semibold">読み込み中...</div>
            </div>
        );
    }

    const userInfo = data.my_info;

    return (
        <div className="min-h-screen bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-3xl mx-auto">
                <div className="bg-white shadow rounded-lg">
                    {/* プロフィールヘッダー */}
                    <div className="px-4 py-5 sm:px-6">
                        <div className="flex items-center">
                            <div className="flex-shrink-0 h-32 w-32">
                                {profileImage ? (
                                    <img
                                        src={profileImage}
                                        alt="プロフィール画像"
                                        className="h-32 w-32 rounded-full object-cover"
                                    />
                                ) : (
                                    <div className="h-32 w-32 rounded-full bg-gray-200 flex items-center justify-center">
                                        <span className="text-gray-500">No Image</span>
                                    </div>
                                )}
                            </div>
                            <div className="ml-6">
                                <h1 className="text-2xl font-bold text-gray-900">
                                    {userInfo.username}
                                </h1>
                                <p className="text-sm text-gray-500 mt-1">{userInfo.email}</p>
                            </div>
                        </div>
                    </div>

                    {/* プロフィール情報 */}
                    <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
                        <dl className="grid grid-cols-1 gap-x-4 gap-y-6 sm:grid-cols-2">
                            <div>
                                <dt className="text-sm font-medium text-gray-500">姓</dt>
                                <dd className="mt-1 text-sm text-gray-900">{userInfo.lastname}</dd>
                            </div>
                            <div>
                                <dt className="text-sm font-medium text-gray-500">名</dt>
                                <dd className="mt-1 text-sm text-gray-900">{userInfo.firstname}</dd>
                            </div>
                            <div>
                                <dt className="text-sm font-medium text-gray-500">生年月日</dt>
                                <dd className="mt-1 text-sm text-gray-900">{new Date(userInfo.birthdate).toLocaleDateString()}</dd>
                            </div>
                            <div>
                                <dt className="text-sm font-medium text-gray-500">性別</dt>
                                <dd className="mt-1 text-sm text-gray-900">{userInfo.gender === 'male' ? '男性' : '女性'}</dd>
                            </div>
                            <div>
                                <dt className="text-sm font-medium text-gray-500">興味のある性別</dt>
                                <dd className="mt-1 text-sm text-gray-900">
                                    {userInfo.sexuality === 'male' ? '男性' :
                                        userInfo.sexuality === 'female' ? '女性' : '両方'}
                                </dd>
                            </div>
                            <div>
                                <dt className="text-sm font-medium text-gray-500">地域</dt>
                                <dd className="mt-1 text-sm text-gray-900">{userInfo.area}</dd>
                            </div>
                            <div className="sm:col-span-2">
                                <dt className="text-sm font-medium text-gray-500">自己紹介</dt>
                                <dd className="mt-1 text-sm text-gray-900">{userInfo.self_intro}</dd>
                            </div>
                            <div className="sm:col-span-2">
                                <dt className="text-sm font-medium text-gray-500">タグ</dt>
                                <dd className="mt-1">
                                    <div className="flex flex-wrap gap-2">
                                        {userInfo.tags.map((tag, index) => (
                                            <span
                                                key={index}
                                                className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
                                            >
                                                {tag}
                                            </span>
                                        ))}
                                    </div>
                                </dd>
                            </div>
                            <div>
                                <dt className="text-sm font-medium text-gray-500">位置情報</dt>
                                <dd className="mt-1 text-sm text-gray-900">
                                    {userInfo.is_gps ? 'GPS使用中' : 'GPS未使用'}
                                </dd>
                            </div>
                            <div>
                                <dt className="text-sm font-medium text-gray-500">評価スコア</dt>
                                <dd className="mt-1 text-sm text-gray-900">{userInfo.fame_rating}</dd>
                            </div>
                        </dl>
                    </div>
                </div>
            </div>
        </div>
    );
}