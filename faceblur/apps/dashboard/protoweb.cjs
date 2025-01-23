const ejs = require('ejs');

const template = `import { customFetch } from "../../../util/fetch";

<%- protocOutput %>
<% services.forEach(service => { %>
<% service.methods.forEach(method => { %>
/*
<%= method.comment %>
*/
export async function <%= method.name %>(request: <%= method.requestType %>, token?: string) {
  return customFetch<<%= method.requestType %>, <%= method.responseType %>>("<%= method.url %>", request, token);
}
<% }); %>
<% }); %>
`;

module.exports = {
    renderTemplate: (data) => {
        return ejs.render(template, data);
    },
};
