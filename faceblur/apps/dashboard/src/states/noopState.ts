import { createContainer } from 'unstated-next';

const noop = () => {};

export const NoopState = createContainer(noop);
