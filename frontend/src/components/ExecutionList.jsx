import { Button, Modal, Popconfirm, Table, Tag, message } from "antd";
import axios from "axios";
import moment from "moment";
import { useCallback, useEffect, useState } from "react";

const ExecutionList = ({
  isExecutionModalVisible,
  setIsExecutionModalVisible,
  onSelect,
}) => {
  const [isLoading, setIsLoading] = useState(false);
  const [savedCodes, setSavedCodes] = useState([]);

  const fetchSavedCodes = async () => {
    setSavedCodes([]);
    setIsLoading(true);
    try {
      const url = `${baseURL}/code/snippets`;
      const response = await axios.get(url);
      const data = response?.data;
      if (data?.status === "success") {
        setSavedCodes(data?.code_snippets);
      }
    } catch (error) {
      message.error("Unable to load saved codes, please try again later");
    } finally {
      setIsLoading(false);
    }
  };

  const deleteCode = async (id) => {
    setIsLoading(true);
    try {
      const url = `${baseURL}/code/snippets/${id}`;
      const response = await axios.delete(url);
      if (response?.status === 204) {
        message.success("deleted successfully");
        await fetchSavedCodes();
      } else throw new Error();
    } catch (error) {
      message.error("Unable to delete, please try again later");
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchSavedCodes();
  }, []);

  const columns = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "Status",
      dataIndex: "status",
      key: "status",
      render: (status) => (
        <Tag color={status === "success" ? "green" : "red"}>{status}</Tag>
      ),
    },
    {
      title: "Timestamp",
      dataIndex: "timestamp",
      key: "timestamp",
      render: (time) => moment(time).format("Do MMM'YY hh:mm A"),
    },
    {
      title: "Actions",
      dataIndex: "id",
      key: "action",
      width: "30%",
      render: (id, data) => (
        <>
          <Popconfirm
            title="Delete code"
            description="Are you sure to delete this code?"
            onConfirm={() => {
              deleteCode(id);
            }}
            okText="Yes"
            cancelText="No"
          >
            <Button className="me-2" danger>
              Delete
            </Button>
          </Popconfirm>
          <Button
            onClick={() => {
              onSelect(data?.code);
              setIsExecutionModalVisible(false);
            }}
          >
            View Code
          </Button>
        </>
      ),
    },
  ];

  const data = savedCodes?.map((data) => ({
    key: data?.id,
    id: data?.id,
    name: data?.name,
    code: data?.code,
    timestamp: data?.date_created,
    status: data?.status,
  }));

  return (
    <Modal
      width="60%"
      title="Saved Codes"
      centered
      open={isExecutionModalVisible}
      onCancel={() => setIsExecutionModalVisible(false)}
      footer={[]}
    >
      {savedCodes?.length > 0 ? (
        <Table
          columns={columns}
          dataSource={data}
          pagination={false}
          loading={isLoading}
        />
      ) : (
        <p>No saved codes</p>
      )}
    </Modal>
  );
};

export default ExecutionList;
