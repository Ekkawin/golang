import { Axios } from '../axios/axios';
interface Props {
  userName: string;
  password: string;
}
export const submitUser = ({ userName, password }: Props) => {
  const response = Axios.get(`/${userName}/${password}`);

  return response;
};
