import clsx from "clsx";
import { FC, ReactNode } from "react";

interface I_ContainerProps {
  borderX?: boolean;
  borderT?: boolean;
  borderB?: boolean;
  className?: string;
  children: ReactNode;
}

export const Container: FC<I_ContainerProps> = (props) => {
  return (
    <div
      className={clsx(
        "container w-full px-4 mx-auto md:px-8 lg:px-12 xl:px-16 2xl:px-20",
        {
          "border-x border-solid": props.borderX,
          "border-b border-solid": props.borderB,
          "border-t border-solid": props.borderT,
        },
        props.className,
      )}
    >
      {props.children}
    </div>
  );
};
