// app/components/UserInfoForm.tsx

import { Form } from "@remix-run/react";
//import type { ActionData } from "~/types/userInfo";
import { useState } from "react";
import { ActionData } from "~/routes/Init-setup";

interface UserInfoFormProps {
    actionData?: ActionData;
    isSubmitting: boolean;
}

export function UserInfoForm({ actionData, isSubmitting }: UserInfoFormProps) {
    const [previewUrl, setPreviewUrl] = useState<string | null>(null);

    const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (file) {
            const reader = new FileReader();
            reader.onloadend = () => {
                setPreviewUrl(reader.result as string);
            };
            reader.readAsDataURL(file);
        }
    };

    return (
        <Form method="post" className="space-y-6" encType="multipart/form-data">
            <div className="grid grid-cols-2 gap-4">
                <div>
                    <label htmlFor="lastname" className="block text-sm font-medium text-gray-700">
                        姓
                    </label>
                    <div className="mt-1">
                        <input
                            id="lastname"
                            name="lastname"
                            type="text"
                            required
                            defaultValue={actionData?.fields?.lastname}
                            className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        />
                        {actionData?.fieldErrors?.lastname && (
                            <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.lastname}</p>
                        )}
                    </div>
                </div>

                <div>
                    <label htmlFor="firstname" className="block text-sm font-medium text-gray-700">
                        名
                    </label>
                    <div className="mt-1">
                        <input
                            id="firstname"
                            name="firstname"
                            type="text"
                            required
                            defaultValue={actionData?.fields?.firstname}
                            className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        />
                        {actionData?.fieldErrors?.firstname && (
                            <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.firstname}</p>
                        )}
                    </div>
                </div>
            </div>

            <div>
                <label htmlFor="birthdate" className="block text-sm font-medium text-gray-700">
                    生年月日
                </label>
                <div className="mt-1">
                    <input
                        id="birthdate"
                        name="birthdate"
                        type="date"
                        required
                        defaultValue={actionData?.fields?.birthdate}
                        className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                    {actionData?.fieldErrors?.birthdate && (
                        <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.birthdate}</p>
                    )}
                </div>
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-700">性別</label>
                <div className="mt-1 space-x-4">
                    <label className="inline-flex items-center">
                        <input
                            type="radio"
                            name="gender"
                            value="male"
                            defaultChecked={actionData?.fields?.gender === "male"}
                            className="form-radio h-4 w-4 text-indigo-600"
                        />
                        <span className="ml-2">男性</span>
                    </label>
                    <label className="inline-flex items-center">
                        <input
                            type="radio"
                            name="gender"
                            value="female"
                            defaultChecked={actionData?.fields?.gender === "female"}
                            className="form-radio h-4 w-4 text-indigo-600"
                        />
                        <span className="ml-2">女性</span>
                    </label>
                    {actionData?.fieldErrors?.gender && (
                        <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.gender}</p>
                    )}
                </div>
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-700">お相手の性別</label>
                <div className="mt-1 space-x-4">
                    <label className="inline-flex items-center">
                        <input
                            type="radio"
                            name="sexuality"
                            value="male"
                            defaultChecked={actionData?.fields?.sexuality === "male"}
                            className="form-radio h-4 w-4 text-indigo-600"
                        />
                        <span className="ml-2">男性</span>
                    </label>
                    <label className="inline-flex items-center">
                        <input
                            type="radio"
                            name="sexuality"
                            value="female"
                            defaultChecked={actionData?.fields?.sexuality === "female"}
                            className="form-radio h-4 w-4 text-indigo-600"
                        />
                        <span className="ml-2">女性</span>
                    </label>
                    <label className="inline-flex items-center">
                        <input
                            type="radio"
                            name="sexuality"
                            value="male/female"
                            defaultChecked={actionData?.fields?.sexuality === "male/female"}
                            className="form-radio h-4 w-4 text-indigo-600"
                        />
                        <span className="ml-2">両方</span>
                    </label>
                    {actionData?.fieldErrors?.sexuality && (
                        <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.sexuality}</p>
                    )}
                </div>
            </div>

            <div>
                <label htmlFor="area" className="block text-sm font-medium text-gray-700">
                    お住まいの地域
                </label>
                <div className="mt-1">
                    <select
                        id="area"
                        name="area"
                        required
                        defaultValue={actionData?.fields?.area}
                        className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    >
                        <option value="">選択してください</option>
                        <option value="Hokkaido">北海道</option>
                        <option value="Aomori">青森県</option>
                        <option value="Iwate">岩手県</option>
                        <option value="Miyagi">宮城県</option>
                        <option value="Akita">秋田県</option>
                        <option value="Yamagata">山形県</option>
                        <option value="Fukushima">福島県</option>
                        <option value="Ibaraki">茨城県</option>
                        <option value="Tochigi">栃木県</option>
                        <option value="Gunma">群馬県</option>
                        <option value="Saitama">埼玉県</option>
                        <option value="Chiba">千葉県</option>
                        <option value="Tokyo">東京都</option>
                        <option value="Kanagawa">神奈川県</option>
                        <option value="Niigata">新潟県</option>
                        <option value="Toyama">富山県</option>
                        <option value="Ishikawa">石川県</option>
                        <option value="Fukui">福井県</option>
                        <option value="Yamanashi">山梨県</option>
                        <option value="Nagano">長野県</option>
                        <option value="Gifu">岐阜県</option>
                        <option value="Shizuoka">静岡県</option>
                        <option value="Aichi">愛知県</option>
                        <option value="Mie">三重県</option>
                        <option value="Shiga">滋賀県</option>
                        <option value="Kyoto">京都府</option>
                        <option value="Osaka">大阪府</option>
                        <option value="Hyogo">兵庫県</option>
                        <option value="Nara">奈良県</option>
                        <option value="Wakayama">和歌山県</option>
                        <option value="Tottori">鳥取県</option>
                        <option value="Shimane">島根県</option>
                        <option value="Okayama">岡山県</option>
                        <option value="Hiroshima">広島県</option>
                        <option value="Yamaguchi">山口県</option>
                        <option value="Tokushima">徳島県</option>
                        <option value="Kagawa">香川県</option>
                        <option value="Ehime">愛媛県</option>
                        <option value="Kochi">高知県</option>
                        <option value="Fukuoka">福岡県</option>
                        <option value="Saga">佐賀県</option>
                        <option value="Nagasaki">長崎県</option>
                        <option value="Kumamoto">熊本県</option>
                        <option value="Oita">大分県</option>
                        <option value="Miyazaki">宮崎県</option>
                        <option value="Kagoshima">鹿児島県</option>
                        <option value="Okinawa">沖縄県</option>
                    </select>
                    {actionData?.fieldErrors?.area && (
                        <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.area}</p>
                    )}
                </div>
            </div>

            <div>
                <label htmlFor="self_intro" className="block text-sm font-medium text-gray-700">
                    自己紹介
                </label>
                <div className="mt-1">
                    <textarea
                        id="self_intro"
                        name="self_intro"
                        rows={4}
                        required
                        defaultValue={actionData?.fields?.selfIntro}
                        className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        placeholder="自己紹介を入力してください（10文字以上1000文字以内）"
                    />
                    {actionData?.fieldErrors?.selfIntro && (
                        <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.selfIntro}</p>
                    )}
                </div>
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-700">プロフィール画像</label>
                <div className="mt-1 flex items-center space-x-4">
                    <div className="flex-shrink-0">
                        {previewUrl ? (
                            <img
                                src={previewUrl}
                                alt="Preview"
                                className="h-24 w-24 rounded-full object-cover"
                            />
                        ) : (
                            <div className="h-24 w-24 rounded-full bg-gray-200 flex items-center justify-center">
                                <span className="text-gray-400">未選択</span>
                            </div>
                        )}
                    </div>
                    <div className="flex-1">
                        <input
                            type="file"
                            name="image"
                            accept="image/jpeg,image/png,image/gif"
                            onChange={handleImageChange}
                            required
                            className="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-indigo-50 file:text-indigo-700 hover:file:bg-indigo-100"
                        />
                        {actionData?.fieldErrors?.image && (
                            <p className="mt-2 text-sm text-red-600">{actionData.fieldErrors.image}</p>
                        )}
                    </div>
                </div>
                <p className="mt-2 text-sm text-gray-500">
                    JPEG、PNG、GIF形式（5MB以下）の画像をアップロードしてください
                </p>
            </div>

            {actionData?.formError && (
                <div className="rounded-md bg-red-50 p-4">
                    <div className="flex">
                        <div className="ml-3">
                            <p className="text-sm font-medium text-red-800">
                                {actionData.formError}
                            </p>
                        </div>
                    </div>
                </div>
            )}

            <div>
                <button
                    type="submit"
                    disabled={isSubmitting}
                    className={`w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 ${isSubmitting ? "opacity-50 cursor-not-allowed" : ""
                        }`}
                >
                    {isSubmitting ? "送信中..." : "プロフィールを設定する"}
                </button>
            </div>
        </Form>
    );
}