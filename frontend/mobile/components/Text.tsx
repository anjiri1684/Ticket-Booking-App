import { defaultShortcuts, ShortcutProps } from "@/styles/shortcuts";
import { PropsWithChildren } from "react";
import { TextProps, Text as RNText} from "react-native";

interface CustomeTextProps extends PropsWithChildren, ShortcutProps, TextProps {
     fontSize?: number;
     bold?: boolean;
     underline?: boolean;
     color?: string
}


export function Text({ fontSize = 18, bold, underline, color, children, ...restProps }: CustomeTextProps) {
     return (
          <RNText style={[defaultShortcuts(restProps), {
               fontSize,
               fontWeight: bold ? 'bold' : "normal",
               textDecorationLine: underline ? 'underline' : 'none',
               color
          }]}{...restProps}>
               { children}
          </RNText>
     )
 }