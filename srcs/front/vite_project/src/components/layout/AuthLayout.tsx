import AuthHeader from "../headers/AuthHeader/AuthHeader.tsx";
import { Outlet } from "npm:react-router-dom";

const AuthLayout: React.FC = () => {
  return (
    <>
      <AuthHeader />
      <Outlet />
    </>
  );
};

export default AuthLayout;
