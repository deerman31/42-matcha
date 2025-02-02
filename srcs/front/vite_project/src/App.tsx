import "./App.css";

import { BrowserRouter as Router, Route, Routes } from "npm:react-router-dom";

import Landing from "./pages/Auth/Landing/Landing.tsx";
import Register from "./pages/Auth/Register/Register.tsx";
import Login from "./pages/Auth/Login/Login.tsx";

import Footer from "./components/footer/Footer.tsx";
import AuthLayout from "./components/layout/AuthLayout.tsx";
import UserLayout from "./components/layout/UserLayout.tsx";

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route element={<AuthLayout />}>
          <Route path="/" element={<Landing />} />
          <Route path="register" element={<Register />} />
          <Route path="login" element={<Login />} />
        </Route>

        <Route element={<UserLayout />}>
        </Route>
      </Routes>

      <Footer />
    </Router>
  );
};

export default App;
