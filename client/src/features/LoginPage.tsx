/** @jsxRuntime classic */
/** @jsx jsx */
import { Button, Form, Input } from 'antd';
import React from 'react';
import { submitUser } from '../api/submitUser';
import { PageWrapper } from './PageWrapper';
import { css, jsx } from '@emotion/react';
import { useHistory } from 'react-router-dom';

const handleSubmit = async (value: any) => {
  console.log(`value`, value);
  const { userName, password } = value;
  const res = await submitUser({ userName, password });
  console.log(`res`, res);
};

export const LoginPage = () => {
  const [form] = Form.useForm();
  const history = useHistory();
  return (
    <PageWrapper
      css={css`
        background-color: #89cff0;
      `}
    >
      <div
        className="bg-white p-12 rounded-2xl"
        css={css`
          width: 512px;
          height: 512px;
        `}
      >
        <Form form={form} onFinish={handleSubmit}>
          <h3 className="text-4xl text-center text-gray-700 font-bold mt-4">
            Log in to Your Account
          </h3>
          <Form.Item name="userName" className="mt-8">
            <Input placeholder="Email" />
          </Form.Item>
          <Form.Item name="password" className="mt-8">
            <Input placeholder="Password" />
          </Form.Item>
          <Button
            className="w-full rounded-lg mt-12 bg-blue-500 border-blue-500 text-white font-bold text-xl leading-none hover:bg-blue-700 hover:border-blue-700"
            htmlType="submit"
          >
            Login
          </Button>
        </Form>
        <div className="text-sm text-gray-400 flex items-center justify-start mt-4">
          Don't have account?
          <span
            className="text-blue-600 ml-2 cursor-pointer hover:text-blue-800"
            onClick={() => {
              console.log('hi');

              history.push('/register');
            }}
          >
            click here
          </span>
        </div>
      </div>
    </PageWrapper>
  );
};
