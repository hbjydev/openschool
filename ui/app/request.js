const { Socket } = require("net");
const { OSPResponse } = require("./response.js");

class OSPRequest {
  /**
   * @param {string} method 
   * @param {string} osrn 
   * @param {string} version 
   * @param {Map<string, string>} headers
   */
  constructor(method, osrn, version, headers = new Map(), body = '') {
    this.method = method;
    this.osrn = osrn;
    this.version = version;
    this.headers = headers;
    this.body = body;
  }

  get headersList() {
    let headers = [];

    this.headers.forEach((v, k) => {
      headers.push(`${k}: ${v}`);
    });

    return headers.join('\r\n');
  }

  /**
   * @returns {string}
   */
  toString() {
    let request = `${this.method} ${this.osrn} ${this.version}`;

    if (this.headers.size > 0) {
      request = `${request}\r\n${this.headersList}`
    }

    if (this.body.size > 0) {
      request = `${request}\r\n\r\n${this.body}`
    }

    return `${request}\r\n\r\n`
  }

  /**
   * @returns {Promise<OSPResponse>}
   */
  send() {
    return new Promise((resolve, reject) => {
      const socket = new Socket();
      socket.connect(8001, "127.0.0.1");

      socket.on("connect", () => {
        socket.write(this.toString());
      });

      socket.on("error", (err) => {
        reject(err);
      });

      socket.on("data", (data) => {
        const response = OSPResponse.fromString(data.toString());
        resolve(response);
      });
    });
  }
}

module.exports = { OSPRequest };
