import clsx from "clsx";
import { FC, ReactNode } from "react";

interface I_SectionProps {
  borderX?: boolean;
  borderB?: boolean;
  borderT?: boolean;
  className?: string;
  children: ReactNode;
}

export const Section: FC<I_SectionProps> = (props) => {
  return (
    <section
      className={clsx(
        "w-full mx-auto",
        {
          "border-x border-solid": props.borderX,
          "border-b border-solid": props.borderB,
          "border-t border-solid": props.borderT,
        },
        props.className,
      )}
    >
      {props.children}
    </section>
  );
};
