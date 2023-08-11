import React, { useState } from "react";
import AceEditor from "react-ace";
import "ace-builds/src-noconflict/mode-golang";
import "ace-builds/src-noconflict/theme-github";
import "ace-builds/src-noconflict/ext-language_tools"

// CodeEditor for golang code
const CodeEditor = ({ code, onChange }) => {
  return (
    <>
      <h6 className="m-3">main.go</h6>
      <hr className="m-0" />
      <AceEditor
        placeholder="Write your golang code here"
        mode="golang"
        theme="github"
        name="blah2"
        onChange={onChange}
        fontSize={14}
        showPrintMargin={true}
        showGutter={true}
        width="100%"
        height="calc(100vh - 200px)"
        highlightActiveLine={true}
        value={code}
        setOptions={{
          enableBasicAutocompletion: true,
          enableLiveAutocompletion: true,
          enableSnippets: true,
          showLineNumbers: true,
          tabSize: 2,
        }}
      />
    </>
  );
};

export default CodeEditor;
