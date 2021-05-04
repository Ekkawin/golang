/** @jsxRuntime classic */
/** @jsx jsx */
import { Button, Form, Input } from 'antd';
import React, { useState } from 'react';
import { createUser } from '../api/createUser';
import { PageWrapper } from './PageWrapper';
import { css, jsx } from '@emotion/react';
import { useHistory } from 'react-router-dom';

export const RegisterPage = () => {
  const [form] = Form.useForm();
  const history = useHistory();
  const [creating, setCreating] = useState(false);

  const onFinish = async (value: any) => {
    const { email, password } = value;
    try {
      setCreating(true);
      await createUser({ email, password });
      history.push('/home');
    } catch (error) {
      console.error(error);
    } finally {
      setCreating(false);
    }
  };
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
        <Form form={form} onFinish={onFinish}>
          <h3 className="text-4xl text-center text-gray-700 font-bold mt-4">
            Create Your Account
          </h3>
          <Form.Item name="email" className="mt-8">
            <Input placeholder="Email" />
          </Form.Item>
          <Form.Item name="password" className="mt-8">
            <Input.Password placeholder="Password" />
          </Form.Item>
          <Form.Item
            name="confirmPassword"
            className="mt-8"
            rules={[
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || getFieldValue('password') === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(
                    new Error(
                      'The two passwords that you entered do not match!'
                    )
                  );
                },
              }),
            ]}
          >
            <Input.Password placeholder="Repeat your Password" />
          </Form.Item>
          <Button
            loading={creating}
            className="w-full rounded-lg mt-12 bg-blue-500 border-blue-500 text-white font-bold text-xl leading-none hover:bg-blue-700 hover:border-blue-700"
            htmlType="submit"
          >
            Login
          </Button>
        </Form>
        <div className="text-sm text-gray-400 flex items-center justify-start mt-4">
          already have account?
          <span className="text-blue-600 ml-2 cursor-pointer hover:text-blue-800">
            sign up
          </span>
        </div>
      </div>
    </PageWrapper>
  );
};
