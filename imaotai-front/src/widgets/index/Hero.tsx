import React from "react";
import { Typography } from "antd";

const { Title } = Typography;
const Hero: React.FC = () => (
  <>
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Title>摸鱼小助手-Imaotai小助手</Title>
    </div>
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <img src="../../../public/logo.svg"></img>
    </div>
  </>
);

export default Hero;
