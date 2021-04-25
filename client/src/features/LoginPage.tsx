import { Button, Form, Input } from 'antd';
import React from 'react';
import { submitUser } from '../api/submitUser';
import { PageWrapper } from './PageWrapper';

const handleSubmit = async (value: any) => {
  console.log(`value`, value);
  const { userName, password } = value;
  const res = await submitUser({ userName, password });
  console.log(`res`, res);
};

export const LoginPage = () => {
  const [form] = Form.useForm();
  return (
    <PageWrapper>
      <Form form={form} onFinish={handleSubmit}>
        <Form.Item label="userName" name="userName">
          <Input />
        </Form.Item>
        <Form.Item label="password" name="password">
          <Input />
        </Form.Item>
        <Button htmlType="submit"> ส่ง</Button>
      </Form>
    </PageWrapper>
  );
};
