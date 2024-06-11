import React, { ReactElement, useEffect, useState } from "react";
import "./css/modal.css";
import {
  Alert,
  Button,
  Form,
  Input,
  Modal,
  Result,
  Space,
  Table,
  Tag,
  Typography,
  message,
} from "antd";
import type { FormProps, TableProps } from "antd";
import {
  GetFollowerUserList,
  DeleteFollowerUser,
  SuspendFollowerUser,
  StartFollowerUser,
  GetVerifyCode,
  UpdateFollowerUserToken,
  UpdateFollowerUserAddress,
} from "../api/user";
interface DataType {
  uid: string;
  key: string;
  mobile: string;
  status: string;
  address: string;
  user_name: string;
  exp_time: string;
}
const columns = (
  setUpdateTokenShow: (cb: (open: boolean) => boolean) => void,
  setUpdateAddressShow: (cb: (open: boolean) => boolean) => void,
  setUid: (uid: string) => void,
  setload: (load: string) => void,
  refresh: () => void
): TableProps<DataType>["columns"] => [
  {
    title: "uid",
    dataIndex: "uid",
    key: "uid",
    render: (text) => <a>{text}</a>,
  },
  {
    title: "手机号",
    dataIndex: "mobile",
    key: "mobile",
  },
  {
    title: "用户名",
    dataIndex: "user_name",
    key: "user_name",
  },
  {
    title: "状态",
    dataIndex: "status",
    key: "status",
  },
  {
    title: "过期时间",
    dataIndex: "exp_time",
    key: "exp_time",
  },
  {
    title: "地址",
    dataIndex: "address",
    key: "address",
  },
  {
    title: "操作",
    key: "action",
    render: (_, record) => (
      <Space size="middle">
        <Button
          type="primary"
          onClick={() => {
            setUid(record.uid);
            setUpdateTokenShow((old) => !old);
          }}
        >
          更新token
        </Button>

        <Button
          type="primary"
          onClick={() => {
            setUid(record.uid);
            setUpdateAddressShow((old) => !old);
          }}
        >
          更新地址
        </Button>
        <Button
          type="primary"
          onClick={() => {
            suspendFollowerUserSubscription(record.uid, setload);
            refresh();
          }}
        >
          暂停申购
        </Button>
        <Button
          type="primary"
          onClick={() => {
            startFollowerUserSubscription(record.uid, setload);
            refresh();
          }}
        >
          继续申购
        </Button>
        <Button
          type="primary"
          onClick={() => {
            deleteFollowerUser(record.uid, setload);
            refresh();
          }}
        >
          删除用户
        </Button>
      </Space>
    ),
  },
];

const getVerifyCode = async (uid: string) => {
  GetVerifyCode({ uid }).then((data) => {
    if (data.data.err_no !== 0) {
      message.error(data.data.err_msg, 2);
      return;
    }
    message.success("发送验证码成功", 2);
  });
};
const suspendFollowerUserSubscription = async (
  uid: string,
  setload: (load: string) => void
) => {
  SuspendFollowerUser({ uid }).then((data) => {
    if (data.data.err_no !== 0) {
      message.error(data.data.err_msg);
      return;
    }
    setload(Math.floor(Date.now() / 1000).toString());
    message.success("暂停申购成功", 2);
  });
};
const startFollowerUserSubscription = async (
  uid: string,
  setload: (load: string) => void
) => {
  StartFollowerUser({ uid }).then((data) => {
    if (data.data.err_no !== 0) {
      message.error(data.data.err_msg);
      return;
    }
    setload(Math.floor(Date.now() / 1000).toString());
    message.success("继续申购成功", 2);
  });
};
const deleteFollowerUser = async (
  uid: string,
  setload: (load: string) => void
) => {
  DeleteFollowerUser({ uid }).then((data) => {
    if (data.data.err_no !== 0) {
      message.error(data.data.err_msg);
      return;
    }
    setload(Math.floor(Date.now() / 1000).toString());
    message.success("删除用户成功", 2);
  });
};

const User: React.FC = () => {
  const [showData, setShowData] = useState<DataType[]>([]);
  const [updateAddressShow, setUpdateAddressShow] = useState<boolean>(false);
  const [updateTokenShow, setUpdateTokenShow] = useState<boolean>(false);
  const [uid, setUid] = useState<string>("");
  const [code, setCode] = useState<string>("");
  const [address, setAddress] = useState<string>("");
  const [load, setLoad] = useState<string>("");
  const getCode = (e: { target: { value: React.SetStateAction<string> } }) => {
    setCode(e.target.value);
  };
  const getAddress = (e: {
    target: { value: React.SetStateAction<string> };
  }) => {
    setAddress(e.target.value);
  };
  const loadData = () => {
    let page = 1;
    let page_size = 25;
    GetFollowerUserList({ page, page_size }).then((data) => {
      let v = data.data.data;
      const newData: DataType[] = [];
      for (let i = 0; i < v.length; i++) {
        newData.push({
          uid: v[i].uid,
          key: i + "",
          exp_time: v[i].exp_time,
          mobile: v[i].mobile,
          status: v[i].status,
          address: v[i].address,
          user_name: v[i].user_name,
        });
      }
      setShowData(newData);
    });
  };

  useEffect(() => {
    loadData();
  }, [load]);

  return (
    <>
      <Table
        columns={columns(
          setUpdateTokenShow,
          setUpdateAddressShow,
          setUid,
          setLoad,
          loadData
        )}
        dataSource={showData}
      />
      <Modal
        open={updateTokenShow}
        footer={<></>}
        onCancel={() => {
          setUpdateTokenShow(false);
          setUid("");
        }}
        closable={false}
        width={500}
        height={700}
        className="user-modal"
      >
        <div style={{ width: "100%", height: "60%" }}>
          <Typography.Title level={5}>验证码</Typography.Title>
          <Input placeholder="imaotai验证码" onChange={getCode} value={code} />
        </div>
        <div
          style={{
            width: "100%",
            height: "40%",
            display: "flex",
            justifyContent: "flex-end",
            alignItems: "center",
            borderTop: "1px solid rgba(233,233,233,1)",
          }}
        >
          <div style={{ width: 160 }}>
            <Button
              type="primary"
              danger
              onClick={() => {
                getVerifyCode(uid);
              }}
              style={{ width: 150, height: 40 }}
            >
              获取验证码
            </Button>
          </div>

          <Button
            type="primary"
            onClick={() => {
              setUpdateTokenShow(false);
              if (code.length !== 6) {
                message.error("请输入6位的验证码后再提交", 2);
                return;
              }
              UpdateFollowerUserToken({ uid, code }).then((data) => {
                if (data.data.err_no !== 0) {
                  message.error(data.data.err_msg, 2);
                  return;
                }
                message.success("更新token成功", 2);
                setCode("");
                setUid("");
                loadData();
                setLoad(Math.floor(Date.now() / 1000).toString());
              });
            }}
            style={{ width: 100, height: 40 }}
          >
            确认
          </Button>
          <div style={{ width: 10 }} />
          <Button
            onClick={() => {
              setUpdateTokenShow(false);
              setUid("");
              setCode("");
            }}
            style={{ width: 100, height: 40 }}
          >
            取消
          </Button>
        </div>
      </Modal>
      <Modal
        open={updateAddressShow}
        footer={<></>}
        onCancel={() => {
          setUpdateAddressShow(false);
          setUid("");
          setCode("");
        }}
        closable={false}
        width={500}
        height={700}
        className="user-modal"
      >
        <div style={{ width: "100%", height: "60%" }}>
          <Typography.Title level={5}>地址</Typography.Title>
          <Input
            placeholder="抢购地址（自己所在地）"
            onChange={getAddress}
            value={address}
          />
        </div>
        <div
          style={{
            width: "100%",
            height: "40%",
            display: "flex",
            justifyContent: "flex-end",
            alignItems: "center",
            borderTop: "1px solid rgba(233,233,233,1)",
          }}
        >
          <Button
            type="primary"
            onClick={() => {
              setUpdateAddressShow(false);
              if (address.length === 0) {
                message.error("请输入地址后再提交", 2);
                return;
              }
              UpdateFollowerUserAddress({ uid, address }).then((data) => {
                if (data.data.err_no !== 0) {
                  message.error(data.data.err_msg, 2);
                  return;
                }
                message.success("更新地址成功", 2);
                setLoad(Math.floor(Date.now() / 1000).toString());
                setAddress("");
                setUid("");
              });
            }}
            style={{ width: 100, height: 40 }}
          >
            确认
          </Button>
          <div style={{ width: 10 }} />
          <Button
            onClick={() => {
              setUpdateAddressShow(false);
              setUid("");
              setAddress("");
            }}
            style={{ width: 100, height: 40 }}
          >
            取消
          </Button>
        </div>
      </Modal>
      <div style={{ width: "100%", height: "100px" }}>
        <Button
          type="primary"
          style={{ cursor: "not-allowed", float: "right" }}
        >
          添加用户(待开发)
        </Button>
      </div>
    </>
  );
};
export default User;
