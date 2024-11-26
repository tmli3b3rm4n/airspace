import React, { useState } from "react";
import { MapContainer, TileLayer, useMapEvents } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import "../App.css"; // For styling

const CordinateForm = ({ lat, lng, onInputChange }) => {
    const [responseMessage, setResponseMessage] = useState("");

    const handleCheckCoordinates = async (e) => {
        e.preventDefault();

        try {
            // Call the API
            const res = await fetch(`http://localhost:8080/restricted-airspace/${lat}/${lng}`);
            const data = await res.json();

            // Check response
            if (data.status === "success" && data.message) {
                const isRestricted = data.message.value;
                setResponseMessage(isRestricted ? "Restricted airspace!" : "Safe to fly.");
            } else {
                setResponseMessage("Unable to determine airspace status.");
            }
        } catch (error) {
            console.error("Error checking coordinates:", error);
            setResponseMessage("An error occurred. Please try again.");
        }
    };


    return (
        <div style={{ position: "relative", height: "100vh", width: "100%" }}>
            {/* Overlay Component */}
            <div className="coordinate-form" style={{}}>
                <h3>Check Coordinates</h3>
                <form
                    onSubmit={(e) => {
                        e.preventDefault();
                        handleCheckCoordinates(e);
                    }}
                >
                    <div>
                        <label>
                            Latitude:
                            <input
                                type="number"
                                value={lat}
                                onChange={(e) => onInputChange("latitude", e.target.value)}
                                placeholder="Enter latitude"
                                required
                            />
                        </label>
                    </div>
                    <div>
                        <label>
                            Longitude:
                            <input
                                type="number"
                                value={lng}
                                onChange={(e) => onInputChange("longitude", e.target.value)}
                                placeholder="Enter longitude"
                                required
                            />
                        </label>
                    </div>
                    <button type="submit" style={{marginLeft: "20px"}}>Check</button>
                </form>
                {responseMessage && <p>{responseMessage}</p>}
            </div>
        </div>
    );
}

export default CordinateForm;
