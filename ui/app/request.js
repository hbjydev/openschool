const { Socket } = require("net");
const { OSPResponse } = require("./response.js");

class OSPRequest {
  /**
   * @param {string} method 
   * @param {string} osrn 
   * @param {string} version 
   */
  constructor(method, osrn, version) {
    this.method = method;
    this.osrn = osrn;
    this.version = version;
  }

  /**
   * @returns {string}
   */
  toString() {
    return `${this.method} ${this.osrn} ${this.version}\r\n\r\n`;
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
