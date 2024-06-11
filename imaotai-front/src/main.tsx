import React from "react";
import ReactDOM from "react-dom/client";
import { router } from "./router";
import { RouterView } from "oh-router-react";
import { Spin } from "antd";
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterView
      router={router}
      splash={
        <div
          style={{
            height: "80vh",
            width: "100vw",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
        >
          <Spin size="large" />
        </div>
      }
    />
  </React.StrictMode>
);
