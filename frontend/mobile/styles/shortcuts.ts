import { DimensionValue } from "react-native";

export interface ShortcutProps {
  // Padding
  p?: number;
  pl?: number;
  pr?: number;
  pt?: number;
  pb?: number;
  py?: number;
  px?: number;

  // Margin
  m?: number | "auto";
  ml?: number | "auto";
  mr?: number | "auto";
  mt?: number | "auto";
  mb?: number | "auto";
  my?: number | "auto";
  mx?: number | "auto";

  // Dimensions
  w?: DimensionValue;
  h?: DimensionValue;
}

export const defaultShortcuts = (props: ShortcutProps) => ({
  // Padding
  padding: props.p,
  paddingLeft: props.pl ?? props.px ?? props.p,
  paddingRight: props.pr ?? props.px ?? props.p,
  paddingTop: props.pt ?? props.py ?? props.p,
  paddingBottom: props.pb ?? props.py ?? props.p,
  paddingVertical: props.py ?? props.p,
  paddingHorizontal: props.px ?? props.p,

  // Margin
  margin: props.m,
  marginLeft: props.ml ?? props.mx ?? props.m,
  marginRight: props.mr ?? props.mx ?? props.m,
  marginTop: props.mt ?? props.my ?? props.m,
  marginBottom: props.mb ?? props.my ?? props.m,
  marginVertical: props.my ?? props.m,
  marginHorizontal: props.mx ?? props.m,

  // Dimensions
  width: props.w,
  height: props.h,
});
