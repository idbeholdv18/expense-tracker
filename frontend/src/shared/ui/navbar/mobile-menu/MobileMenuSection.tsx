import clsx from "clsx";
import { Link } from "./types";
import { FC } from "react";

interface I_MobileMenuSectionProps {
  links: Link[];
}

export const MobileMenuSection: FC<I_MobileMenuSectionProps> = (props) => {
  return (
    <div
      className={clsx(
        "flex flex-col gap-2 bg-amber-300 rounded-sm border border-solid border-border p-4",
      )}
    >
      {props.links.map((link, index) => (
        <p key={index}>{link.name}</p>
      ))}
    </div>
  );
};
