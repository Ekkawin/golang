import classNames from 'classnames';
export const PageWrapper = (props: any) => {
  const { children, className, ...restProps } = props;
  return (
    <div
      className={classNames(
        'flex justify-center items-center px-4 py-4 w-full min-h-screen',
        className
      )}
      {...restProps}
    >
      {props.children}
    </div>
  );
};
