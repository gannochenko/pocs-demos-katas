import React, { FC } from 'react';

import {AboutPageRoot} from './style';
import { AboutPageProps } from './type';
import {Link, PageLayout} from "../../components";
import {Container} from "@mui/joy";

export const AboutPage: FC<AboutPageProps> = ({
    children,
    ...restProps
}) => {
    return (
        <PageLayout title="About" displayPageTitle>
            <Container>
                <AboutPageRoot {...restProps}>
                    <p>
                        This is a Golang + React demo cloud-native app. No bullshit, only production-tested code.
                    </p>
                    <p>
                        It's based on (but does not strictly follow) <Link href="https://github.com/swagger-api/swagger-petstore">the Swagger-petstore demo API</Link>.
                    </p>
                    <p>
                        Under the hood:
                    </p>
                    <ul>
                        <li>
                            Backend
                            <ul>
                                <li>✅ Golang</li>
                                <li>✅ GORM</li>
                                <li>✅ Dependency injection</li>
                                <li>✅ Error handling</li>
                                <li>✅ Logging</li>
                                <li>✅ Unit/integration testing (examples only)</li>
                                <li>✅ Data fixture generator</li>
                                <li>✅ Docker</li>
                                <li>✅ Authentication with Auth0</li>
                                <li>❌ CICD</li>
                                <li>❌ S3 / GCS</li>
                                <li>❌ O11y (metrics, Prometheus, Grafana)</li>
                            </ul>
                        </li>
                        <li>
                            Frontend
                            <ul>
                                <li>✅ TypeScript</li>
                                <li>✅ React</li>
                                <li>✅ create-react-app</li>
                                <li>✅ react router</li>
                                <li>✅ mui/joy</li>
                                <li>✅ react-hooks</li>
                                <li>✅ unstated-next</li>
                                <li>✅ Authentication with Auth0</li>
                                <li>❌ Unit testing</li>
                            </ul>
                        </li>
                    </ul>
                </AboutPageRoot>
            </Container>
        </PageLayout>
    )
};
