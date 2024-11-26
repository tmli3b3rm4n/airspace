import logo from './logo.svg';
import Leaflet from './components/leaflet'
import './App.css';
import CordinateForm from "./components/CordinateForm";
import React, { useState } from "react";


function App() {
    const [latitude, setLatitude] = useState("");
    const [longitude, setLongitude] = useState("");
    const handleMapClick = (lat, lng) => {
        setLatitude(lat);
        setLongitude(lng);
    };

    const handleInputChange = (field, value) => {
        if (field === "latitude") setLatitude(value);
        if (field === "longitude") setLongitude(value);
    };

    return (
    <div className="App">
      <header className="App-header">
          <Leaflet onMapClick={handleMapClick}>
          </Leaflet>
          <CordinateForm
              lat={latitude}
              lng={longitude}
              onInputChange={handleInputChange}
          />
      </header>
    </div>
  );
}

export default App;
