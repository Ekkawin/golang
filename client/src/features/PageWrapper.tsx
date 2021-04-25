import React from 'react';

export const PageWrapper = (props: any) => {
  return (
    <div className="flex justify-center items-center px-4 py-4 w-full h-full">
      {props.children}
    </div>
  );
};
