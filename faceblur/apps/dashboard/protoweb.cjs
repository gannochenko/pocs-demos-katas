const ejs = require('ejs');

const template = `<% if(services.length) { %>import { customFetch, isErrorResponse } from "../../../util/fetch";
import {ErrorResponse} from "@/proto/common/error/v1/error";<% } %>

<%- protocOutput %>
<% if(hasAnyServices) { %>
type DeepNonUndefined<T> = T extends Date
    ? T
    : T extends object
        ? { [K in keyof T]: DeepNonUndefined<Exclude<T[K], undefined>> }
        : T;

type DeepReplaceDateWithNullable<T> = T extends Date
  ? Date | null
  : T extends object
  ? {
      [K in keyof T]: DeepReplaceDateWithNullable<T[K]>;
    }
  : T;
<% } %>

<% services.forEach(service => { %>
<% service.methods.forEach(method => { %>
type <%= method.responseType %>Strict = DeepNonUndefined<DeepReplaceDateWithNullable<<%= method.responseType %>>>;
/*
<%= method.comment %>
*/
export async function <%= method.name %>(request: <%= method.requestType %>, token?: string): Promise<<%= method.responseType %>Strict | ErrorResponse> {
  const data = await customFetch<<%= method.requestType %>, <%= method.responseType %>>("<%= method.url %>", request, token);
  if (!isErrorResponse(data)) {
    await <%= method.responseType %>Decoder.decodeToPromise(data);
  }

  return data as <%= method.responseType %>Strict | ErrorResponse;
}
<% }); %>
<% }); %>
`;

module.exports = {
    renderTemplate: (data) => {
        let hasAnyServices = false;
        for(let i = 0; i < data.services.length; i++) {
            if (data.services[i].methods.length) {
                hasAnyServices = true;
                break;
            }
        }

        return ejs.render(template, {
            ...data,
            hasAnyServices,
        });
    },
};
