
interface Param {
  type: string;
  operator: string;
  key: string;
  value: ParamValue;
}

interface ParamValue {
  item: string;
  list: string[];
}

interface Form {
  startTime: Date;
  endTime: any;
  namespace: string[];
  source: string[];
  traceID: string[];
  host: string[];
  level: string[];
  buildCommit: string;
  configHash: string;
  message: string;
}


export { Param, Form, ParamValue }
