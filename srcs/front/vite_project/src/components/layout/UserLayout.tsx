import UserHeader from "../headers/UserHeader.tsx";
import { Outlet } from "npm:react-router-dom";

const UserLayout: React.FC = () => {
  return (
    <>
      <UserHeader />
      <Outlet />
    </>
  );
};

export default UserLayout;
