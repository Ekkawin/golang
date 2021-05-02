import { Axios } from '../axios/axios';
interface Props {
  email: string;
  password: string;
}
export const createUser = ({ email, password }: Props) => {
  console.log(`email`, email);
  console.log(`password`, password);
  //   const response = Axios.post(`/user`, {
  //     email,
  //     password,
  //   });

  // return response;
  return 'asdf';
};
