// app/routes/app/_layout.tsx
import { Outlet } from "@remix-run/react";
import { AppHeader } from "~/components/AppHeader";

export default function AppLayout() {
  return (
    <div>
      <AppHeader />
      <main className="p-6">
        <Outlet />
      </main>
    </div>
  );
}