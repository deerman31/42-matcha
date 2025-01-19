// app/routes/logout.tsx
import { ActionFunctionArgs, redirect } from "@remix-run/node";
import { destroySession, getAccessToken } from "~/session.server";

export const action = async ({ request }: ActionFunctionArgs) => {
    const token = await getAccessToken(request);

    if (token) {
        try {
            const response = await fetch("http://back:3000/api/logout", {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });

            if (!response.ok) {
                throw new Error('Logout failed');
            }
        } catch (error) {
            console.error('Error during logout:', error);
        }
    }

    await destroySession(request);
    return redirect("/login");
};