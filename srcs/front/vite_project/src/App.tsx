import "./App.css";

import { BrowserRouter as Router, Route, Routes } from "npm:react-router-dom";

import Landing from "./pages/Auth/Landing/Landing.tsx";
import Register from "./pages/Auth/Register/Register.tsx";
import Login from "./pages/Auth/Login/Login.tsx";

import SetupUserInfo from "./pages/Auth/SetupUserInfo/SetupUserInfo.tsx";

import Footer from "./components/footer/Footer.tsx";
import AuthLayout from "./components/layout/AuthLayout.tsx";
import UserLayout from "./components/layout/UserLayout.tsx";

import MyProfile from "./pages/User/MyProfile/MyProfile.tsx";
import Setting from "./pages/User/Setting/Setting.tsx";
import ChangeUsername from "./pages/User/Setting/ChangeUsername.tsx";

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route element={<AuthLayout />}>
          <Route path="/" element={<Landing />} />
          <Route path="register" element={<Register />} />
          <Route path="login" element={<Login />} />

          <Route path="setup-user-info" element={<SetupUserInfo />} />
        </Route>

        <Route element={<UserLayout />}>
          <Route path="/my-profile" element={<MyProfile />} />
          <Route path="/setting" element={<Setting />} />
          <Route path="/change-username" element={<ChangeUsername />} />
        </Route>
      </Routes>

      <Footer />
    </Router>
  );
};

export default App;
