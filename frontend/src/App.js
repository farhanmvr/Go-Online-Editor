import { useState } from "react";
import "./App.css";
import CodeEditor from "./components/CodeEditor";
import ExecutionList from "./components/ExecutionList";
import { Button, Input, Modal, message } from "antd";
import OutputSection from "./components/OutputSection";
import { baseURL, defaultGoCode } from "./constants";
import axios from "axios";

const App = () => {
  const [code, setCode] = useState(defaultGoCode);
  const [execResult, setExecResult] = useState({});
  const [isExecutionModalVisible, setIsExecutionModalVisible] = useState(false);
  const [isNameModalVisible, setIsNameModalVisible] = useState(false);
  const [isExecuting, setIsExecuting] = useState(false);
  const [isSaving, setIsSaving] = useState(false);
  const [name, setName] = useState("");

  // Run code
  const executeCode = async () => {
    setExecResult({});
    setIsExecuting(true);
    try {
      const url = `${baseURL}/code/execute`;
      const body = {
        code,
      };
      const response = await axios.post(url, body);
      const data = response?.data;
      setExecResult(data);
    } catch (error) {
      setExecResult({ error: "Unable to compile the code, please try again" });
    } finally {
      setIsExecuting(false);
    }
  };

  // Save code to db
  const saveCode = async () => {
    setExecResult({});
    setIsSaving(true);
    try {
      const url = `${baseURL}/code/save`;
      const body = {
        code,
        name,
      };
      const response = await axios.post(url, body);
      const data = response?.data;
      if (data?.status === "success")
        message.success("code saved successfully");
      else
        message.warning(
          "not able to compile your code, please fix the code and try again"
        );
      setExecResult({ error: data?.error });
    } catch (error) {
      setExecResult({ error: "Unable to compile the code, please try again" });
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <div className="app">
      {/* Header button section */}
      <div className="d-flex justify-content-end p-3">
        <Button
          onClick={() => setIsExecutionModalVisible(true)}
          size="large"
          className="me-3"
        >
          Saved Codes
        </Button>
        <Button
          onClick={() => setIsNameModalVisible(true)}
          type="primary"
          size="large"
          className="px-4 me-3"
          loading={isSaving}
        >
          Run & Save
        </Button>
        <Button
          onClick={executeCode}
          type="primary"
          size="large"
          className="px-4"
          loading={isExecuting}
        >
          Run
        </Button>
      </div>
      <div className="row">
        <div className="col code-editor pe-0">
          {/* Golang code Editor */}
          <CodeEditor code={code} onChange={setCode} />
        </div>
        <div className="col code-output px-0">
          {/* Golang code output area */}
          <OutputSection execResult={execResult} />
        </div>
      </div>
      {/* Execution list modal */}
      {isExecutionModalVisible && (
        <ExecutionList
          onSelect={(val) => {
            setExecResult({});
            setCode(val);
          }}
          isExecutionModalVisible={isExecutionModalVisible}
          setIsExecutionModalVisible={setIsExecutionModalVisible}
        />
      )}
      {/* Popup for providing name while saving a code */}
      <Modal
        title="Enter a name to your code"
        open={isNameModalVisible}
        onCancel={() => setIsNameModalVisible(false)}
        onOk={() => {
          if (name?.length < 3)
            message.info("Enter atleast 3 character long name");
          else {
            saveCode();
            setIsNameModalVisible(false);
          }
        }}
      >
        <Input
          placeholder="Enter name"
          value={name}
          onChange={(e) => setName(e?.target?.value)}
        />
      </Modal>
    </div>
  );
};

export default App;
