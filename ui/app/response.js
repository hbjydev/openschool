/**
 * @class
 * @constructor
 */
class OSPResponse {
    constructor() {
        /** @type {string} */
        this.version = null;

        /** @type {number} */
        this.status = null;

        /** @type {string} */
        this.reason = null;

        /** @type {Map<string, string>} */
        this.headers = new Map();

        /** @type {string?} */
        this.body = null;
    }

    /**
     * @param {string} raw 
     * @returns {OSPResponse} response as a class
     */
    static fromString(raw) {
        const response = new OSPResponse();
        const lines = raw.split('\n');

        let isBody = false;

        for (const index in lines) {
            const line = lines[index];

            if (line.length == 0) {
                isBody = true;
                continue;
            }

            if (index == "0") {
                const [version, status, ...reasons] = line.split(' ');
                const reason = reasons.join(' ');

                response.version = version;
                response.status = parseInt(status);
                response.reason = reason;
            } else {
                if (isBody) {
                    if (response.body == undefined) response.body = line;
                    else response.body = `${response.body}\n${line}`;
                } else {
                    const [key, value] = line.split(':');
                    response.headers.set(key, value.trim());
                }
            }
        }

        return response;
    }
}

module.exports = { OSPResponse };
