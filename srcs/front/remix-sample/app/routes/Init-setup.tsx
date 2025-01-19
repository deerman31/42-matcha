// app/routes/init_setup.tsx

// app/routes/user-info.tsx

import type { ActionFunction } from "@remix-run/node";
import { Form, useActionData, useNavigation } from "@remix-run/react";
import { useState } from "react";
import { UserInfoForm } from "~/components/UserInfoForm";
import { redirect } from "@remix-run/node"; // redirectをインポート
import {
  validateName,
  validateBirthdate,
  validateGender,
  validateSexuality,
  validateArea,
  validateSelfIntro,
  validateImage,
} from "../utils/validation";
import { destroySession, getAccessToken } from "~/session.server";

export interface ActionData {
  formError?: string;
  fieldErrors?: {
    lastname?: string;
    firstname?: string;
    birthdate?: string;
    gender?: string;
    sexuality?: string;
    area?: string;
    selfIntro?: string;
    image?: string;
  };
  fields?: {
    lastname: string;
    firstname: string;
    birthdate: string;
    gender: string;
    sexuality: string;
    area: string;
    selfIntro: string;
  };
}

export interface FormData {
  lastname: string;
  firstname: string;
  birthdate: string;
  gender: string;
  sexuality: string;
  area: string;
  selfIntro: string;
  image: File;
}


export const action: ActionFunction = async ({ request }) => {
  // セッションからトークンを取得
  const token = await getAccessToken(request);
  if (!token) {
    return redirect("/login");
  }

  const formData = await request.formData();

  const lastname = formData.get("lastname") as string;
  const firstname = formData.get("firstname") as string;
  const birthdate = formData.get("birthdate") as string;
  const gender = formData.get("gender") as string;
  const sexuality = formData.get("sexuality") as string;
  const area = formData.get("area") as string;
  const selfIntro = formData.get("self_intro") as string;
  const image = formData.get("image") as File;


  if (!lastname || !firstname || !birthdate || !gender || !sexuality || !area || !selfIntro || !image) {
    return new Response(
      JSON.stringify({ formError: "すべての項目を入力してください。" }),
      {
        status: 400,
        headers: { "Content-Type": "application/json" },
      }
    );
  }

  const fields = { lastname, firstname, birthdate, gender, sexuality, area, selfIntro };
  const fieldErrors = {
    lastname: validateName(lastname),
    firstname: validateName(firstname),
    birthdate: validateBirthdate(birthdate),
    gender: validateGender(gender),
    sexuality: validateSexuality(sexuality),
    area: validateArea(area),
    selfIntro: validateSelfIntro(selfIntro),
    image: validateImage(image),
  };

  if (Object.values(fieldErrors).some(Boolean)) {
    return new Response(
      JSON.stringify({ fieldErrors, fields }),
      {
        status: 400,
        headers: { "Content-Type": "application/json" },
      }
    );
  }

  try {
    const formDataToSend = new FormData();
    formDataToSend.append("lastname", lastname);
    formDataToSend.append("firstname", firstname);
    formDataToSend.append("birthdate", birthdate);
    formDataToSend.append("gender", gender);
    formDataToSend.append("sexuality", sexuality);
    formDataToSend.append("area", area);
    formDataToSend.append("self_intro", selfIntro);
    formDataToSend.append("image", image);


    const response = await fetch("http://back:3000/api/users/set/user-info", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: formDataToSend,
    });


    // レスポンスのJSONデータを取得
    //const responseData = await response.json();

    if (!response.ok) {
      if (response.status === 401) {
        destroySession(request);
        return redirect("/login");
      }
      const errorData = await response.json();
      return new Response(
        JSON.stringify({ formError: errorData.error || "更新エラー。" }),
        {
          status: response.status,
          headers: { "Content-Type": "application/json" },
        }
      );
    }

    return redirect("/my-profile");
  } catch (error) {
    return new Response(
      JSON.stringify({
        formError: "プロフィール更新中にエラーが発生しました。",
        fields,
      }),
      {
        status: 500,
        headers: { "Content-Type": "application/json" },
      }
    );
  }
};

export default function UserInfoPage() {
  const actionData = useActionData<ActionData>();
  const navigation = useNavigation();
  const isSubmitting = navigation.state === "submitting";

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
          プロフィール設定
        </h2>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          <UserInfoForm
            actionData={actionData}
            isSubmitting={isSubmitting}
          />
        </div>
      </div>
    </div>
  );
}