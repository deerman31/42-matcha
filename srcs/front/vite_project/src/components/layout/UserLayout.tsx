import UserHeader from "../headers/UserHeader.tsx";
import { Outlet } from "npm:react-router-dom";
import LocationService from "../LocationService/LocationService.tsx";

const UserLayout: React.FC = () => {
  return (
    <>
      <UserHeader />
      <LocationService />
      <Outlet />
    </>
  );
};

export default UserLayout;
