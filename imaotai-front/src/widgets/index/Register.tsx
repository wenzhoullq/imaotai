import React from "react";
import type { FormProps } from "antd";
import { Button, Checkbox, Form, Input, message } from "antd";
import { register } from "../../api/index";
import { router } from "../../router";
import { setToken } from "../../shared/token";
import "./css/form.css";
type RegisterType = {
  mobile?: string;
  pass_word?: string;
  email?: string;
  address?: String;
};

const onFinish: FormProps<RegisterType>["onFinish"] = async (values) => {
  var mobile = String(values.mobile);
  var pass_word = String(values.pass_word);
  var address = String(values.address);
  var email = String(values.email);
  await register({ mobile, pass_word, address, email }).then((data) => {
    if (data.data.err_no !== 0) {
      message.error(data.data.err_msg);
      return;
    }
    setToken(data.headers.authorization);
    message.info("注册成功");
    router.navigate("/");
  });
};

const onFinishFailed: FormProps<RegisterType>["onFinishFailed"] = (
  errorInfo
) => {
  console.log("Failed:", errorInfo);
};

const Register: React.FC = () => (
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
      layout="horizontal"
      style={{
        width: 900,
        maxWidth: 600,
      }}
      initialValues={{ remember: true }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <Form.Item<RegisterType>
        label="手机号"
        name="mobile"
        rules={[{ required: true, message: "手机号不能为空" }]}
      >
        <Input />
      </Form.Item>
      <Form.Item<RegisterType>
        label="密码"
        name="pass_word"
        rules={[{ required: true, message: "密码不能为空" }]}
      >
        <Input.Password />
      </Form.Item>
      <Form.Item<RegisterType>
        label="地址"
        name="address"
        rules={[{ required: true, message: "地址不能为空" }]}
      >
        <Input />
      </Form.Item>
      <Form.Item<RegisterType>
        label="邮箱"
        name="email"
        rules={[{ required: true, message: "邮箱不能为空" }]}
      >
        <Input />
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
            注册
          </Button>
        </Form.Item>
      </div>
    </Form>
  </div>
);

export default Register;
