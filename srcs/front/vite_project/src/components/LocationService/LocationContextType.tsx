// LocationContext.tsx
import React, { createContext, useContext, useState } from "react";

interface LocationContextType {
  isTrackingEnabled: boolean;
  setTrackingEnabled: (enabled: boolean) => void;
}

const LocationContext = createContext<LocationContextType | undefined>(
  undefined,
);

interface LocationProviderProps {
  children: React.ReactNode;
}

export const LocationProvider: React.FC<LocationProviderProps> = (
  { children }: LocationProviderProps,
) => {
  const [isTrackingEnabled, setTrackingEnabled] = useState(false);

  return (
    <LocationContext.Provider value={{ isTrackingEnabled, setTrackingEnabled }}>
      {children}
    </LocationContext.Provider>
  );
};

export const useLocationContext = () => {
  const context = useContext(LocationContext);
  if (context === undefined) {
    throw new Error(
      "useLocationContext must be used within a LocationProvider",
    );
  }
  return context;
};
