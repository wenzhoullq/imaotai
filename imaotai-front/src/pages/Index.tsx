import React, { useState } from "react";
import Header from "../widgets/index/Header";
import Hero from "../widgets/index/Hero";
import Login from "../widgets/index/Login";
import Register from "../widgets/index/Register";

const Index: React.FC = () => {
  const [currentForm, setCurrentForm] = useState("login");
  return (
    <div className="container mx-auto p-4">
      <Hero />
      <Header setCurrentForm={setCurrentForm} />
      {currentForm === "login" ? (
        <Login />
      ) : currentForm === "register" ? (
        <Register />
      ) : (
        ""
      )}
    </div>
  );
};

export default Index;
