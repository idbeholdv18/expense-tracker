import { InputHTMLAttributes } from "react";
import { Field } from "./Field";
import { Label } from "./Label";
import { Root } from "./Root";
import { Error } from "./Error";

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
}

export const Input = Object.assign(Root, {
  Field: Field,
  Label: Label,
  Error: Error,
});
