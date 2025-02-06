// NavigationPage.tsx
import React from "react";
import "./Setting.css";
import ChangeUsername from "./UserName/ChangeUsername.tsx";
import ChangeEmail from "./Email/ChangeEmail.tsx";
import ChangeLastname from "./LastName/ChangeLastname.tsx";
import ChangeFirstname from "./FirstName/ChangeFirstname.tsx";
import ChangeBirthDate from "./BirthDate/ChangeBirthDate.tsx";
import ChangeGender from "./Gender/ChangeGender.tsx";
import ChangeSexuality from "./Sexuality/ChangeSexuality.tsx";
import ChangeArea from "./Area/ChangeArea.tsx";
import ChangeIsGps from "./IsGps/IsGps.tsx";
import InteractiveLocationMap from "./Map/InteractiveLocationMap.tsx";
import ChangeSelfIntro from "./SelfIntra/ChangeSelfIntro.tsx";
import TagSearchComponent from "./tag/TagSearchComponent.tsx";

const Setting: React.FC = () => {
  return (
    <div className="navigation-container">
      <div className="button-container">
        <ChangeUsername />
        <ChangeEmail />
        <ChangeLastname />
        <ChangeFirstname />
        <ChangeBirthDate />
        <ChangeGender />
        <ChangeSexuality />
        <ChangeArea />
        <ChangeIsGps />
        <InteractiveLocationMap />
        <ChangeSelfIntro />
        <TagSearchComponent />
      </div>
    </div>
  );
};

export default Setting;
