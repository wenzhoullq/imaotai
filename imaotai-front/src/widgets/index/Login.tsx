import React from "react";
import type { FormProps } from "antd";
import { Button, Checkbox, Form, Input, message } from "antd";
import { login } from "../../api/index";
import { router } from "../../router";
import { setToken } from "../../shared/token";
import "./css/form.css";
type LoginType = {
  mobile?: string;
  pass_word?: string;
};

const onFinish: FormProps<LoginType>["onFinish"] = async (values) => {
  var mobile = String(values.mobile);
  var pass_word = String(values.pass_word);
  await login({ mobile, pass_word }).then((data) => {
    if (data.data.err_no !== 0) {
      message.error(data.data.err_msg);
      return;
    }
    setToken(data.headers.authorization);
    message.info("登录成功");
    router.navigate("/");
  });
};

const onFinishFailed: FormProps<LoginType>["onFinishFailed"] = (errorInfo) => {
  console.log("Failed:", errorInfo);
};

const Login: React.FC = () => (
  <div
    style={{
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    }}
  >
    <Form
      name="basic"
      labelCol={{ span: 5 }}
      wrapperCol={{ span: 16 }}
      style={{ maxWidth: 600, width: 1000 }}
      initialValues={{ remember: true }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <Form.Item<LoginType>
        label="手机号"
        name="mobile"
        rules={[{ required: true, message: "手机号不能为空" }]}
      >
        <Input />
      </Form.Item>

      <Form.Item<LoginType>
        label="密码"
        name="pass_word"
        rules={[{ required: true, message: "密码不能为空" }]}
      >
        <Input.Password />
      </Form.Item>
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button type="primary" htmlType="submit">
            登录
          </Button>
        </Form.Item>
      </div>
    </Form>
  </div>
);

export default Login;
