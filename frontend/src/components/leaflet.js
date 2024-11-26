import React, { useState, useEffect } from "react";
import { MapContainer, TileLayer, useMapEvents, GeoJSON } from "react-leaflet";
import "leaflet/dist/leaflet.css";

const Leaflet = ({ children, onMapClick }) => {
  const [geoJsonData, setGeoJsonData] = useState(null);
    const MapClickHandler = () => {
        useMapEvents({
            click: (event) => {
                const { lat, lng } = event.latlng;
                onMapClick(lat, lng); // Pass clicked coordinates to parent
            },
        });
        return null;
    };
    useEffect(() => {
    // Load the GeoJSON file from the public folder
    fetch("/National_Security_UAS_Flight_Restrictions.geojson")
        .then((response) => {
          if (!response.ok) throw new Error("Failed to load GeoJSON file.");
          return response.json();
        })
        .then((data) => setGeoJsonData(data))
        .catch((error) => console.error("Error loading GeoJSON:", error));
  }, []);

  return (
      <MapContainer center={[37.7749, -122.4194]} zoom={5} style={{ height: "100vh", width: "100%" }}>
          <TileLayer
              url="https://{s}.basemaps.cartocdn.com/dark_all/{z}/{x}/{y}{r}.png"
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/">CARTO</a>'
          />
          <MapClickHandler />
        {geoJsonData && (
            <GeoJSON
                data={geoJsonData}
                style={() => ({
                  color: "red",
                  weight: 2,
                })}
            />
        )}
      </MapContainer>
  );
}

export default Leaflet;