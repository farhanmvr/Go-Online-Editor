import React from "react";

const OutputSection = ({ execResult }) => {
  return (
    <div className="output-section">
      <h6 className="m-3">Output</h6>
      <hr className="m-0" />
      <div className="p-3">
        {execResult?.output && <pre>{execResult?.output}</pre>}
        {execResult?.error && <pre>{execResult?.error}</pre>}
      </div>
    </div>
  );
};

export default OutputSection;
