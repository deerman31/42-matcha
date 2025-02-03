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
import ChangeUsername from "./pages/User/Setting/UserName/ChangeUsername.tsx";
import ChangeEmail from "./pages/User/Setting/Email/ChangeEmail.tsx";
import ChangeLastname from "./pages/User/Setting/LastName/ChangeLastname.tsx";
import ChangeFirstname from "./pages/User/Setting/FirstName/ChangeFirstname.tsx";
import ChangeBirthDate from "./pages/User/Setting/BirthDate/ChangeBirthDate.tsx";
import ChangeGender from "./pages/User/Setting/Gender/ChangeGender.tsx";
import ChangeSexuality from "./pages/User/Setting/Sexuality/ChangeSexuality.tsx";
import ChangeArea from "./pages/User/Setting/Area/ChangeArea.tsx";

import ChangeIsGps from "./pages/User/Setting/IsGps/IsGps.tsx";
import InteractiveLocationMap from "./pages/User/Setting/Map/InteractiveLocationMap.tsx";
import ChangeSelfIntro from "./pages/User/Setting/SelfIntra/ChangeSelfIntro.tsx";


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
          <Route path="/change-email" element={<ChangeEmail />} />
          <Route path="/change-lastname" element={<ChangeLastname />} />
          <Route path="/change-firstname" element={<ChangeFirstname />} />
          <Route path="/change-birthdate" element={<ChangeBirthDate />} />
          <Route path="/change-gender" element={<ChangeGender />} />
          <Route path="/change-sexuality" element={<ChangeSexuality />} />
          <Route path="/change-area" element={<ChangeArea />} />
          <Route path="/change-is-gps" element={<ChangeIsGps />} />
          <Route path="/change-map" element={<InteractiveLocationMap />} />
          <Route path="/change-self-intro" element={<ChangeSelfIntro />} />

        </Route>
      </Routes>

      <Footer />
    </Router>

  );
};

export default App;
