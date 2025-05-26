import { defaultShortcuts, ShortcutProps } from "@/styles/shortcuts";
import { PropsWithChildren } from "react";
import { ViewProps,View } from "react-native";

export interface StackProps extends PropsWithChildren, ShortcutProps, ViewProps {
  flex?: number;
  direction?: "row" | "column";
  gap?: number;
  alignItems?: "flex-start" | "flex-end" | "center" | "stretch" | "baseline";
  justifyContent?:
    | "flex-start"
    | "flex-end"
    | "center"
    | "space-around"
    | "space-between"
    | "space-evenly";
}

export function Stack({
     flex,
     direction,
     gap,
     alignItems,
     justifyContent,
     children,
     style,
     ...restProps
}: StackProps) {
     return (
       <View
         style={[
           defaultShortcuts(restProps),
           {
             flex,
             flexDirection: direction,
             gap,
             alignItems,
             justifyContent,
                    },
           style
         ]}{...restProps}
       >
         {children}
       </View>
     );
}