import { Links, Meta, Scripts, Outlet, useMatches } from "@remix-run/react";
import "./tailwind.css";
import PublicHeader from "./components/PublicHeader";
import { AppHeader } from "./components/AppHeader";


export default function App() {
  const matches = useMatches();
  const currentPath = matches[matches.length - 1]?.pathname || '';
  const isSignupRoute = currentPath.startsWith("/signup");
  const isLoginRoute = currentPath.startsWith("/login");
  const isRootRoute = currentPath === "/";
  return (
    <html lang="ja">
      <head>
        <meta charSet="utf-8" />
        <Meta />
        <Links />
      </head>
      <body>
        {(isSignupRoute || isLoginRoute || isRootRoute) ? <PublicHeader /> : <AppHeader />}
        <main className="p-6">
          <Outlet />
        </main>
        <Scripts />
      </body>
    </html>
  );
}