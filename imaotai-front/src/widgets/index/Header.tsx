import { useState } from "react";
import { Radio, RadioChangeEvent } from "antd";

const Header = (props: { setCurrentForm: (val: string) => void }) => {
  const [radioValue, setRadioValue] = useState<string>("login");
  const onChange = (e: RadioChangeEvent) => {
    props.setCurrentForm(e.target.value);
    setRadioValue(e.target.value);
  };
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "50px",
      }}
    >
      <Radio.Group onChange={onChange} value={radioValue}>
        <Radio value={"login"}>登录</Radio>
        <Radio value={"register"}>注册</Radio>
      </Radio.Group>
    </div>
  );
};

export default Header;
