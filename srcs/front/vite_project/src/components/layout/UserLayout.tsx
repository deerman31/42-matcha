import UserHeader from "../headers/UserHeader.tsx";
import { Outlet } from "npm:react-router-dom";
import LocationService from "../LocationService/LocationService.tsx";
import { LocationProvider } from "../LocationService/LocationContextType.tsx";

const UserLayout: React.FC = () => {
  return (
    <LocationProvider>
      <UserHeader />
      <LocationService />
      <Outlet />
    </LocationProvider>
  );
};

export default UserLayout;
