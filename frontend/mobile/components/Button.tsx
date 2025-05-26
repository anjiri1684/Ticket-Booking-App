import { defaultShortcuts, ShortcutProps } from "@/styles/shortcuts";
import { ActivityIndicator, StyleSheet, TouchableOpacity, TouchableOpacityProps } from "react-native";
import { Text } from "./Text";


interface ButtonProps extends ShortcutProps, TouchableOpacityProps {
     variant?: "contained" | "outlined" | "ghost"
     isLoading?: boolean
}


export function Button({ onPress, children, variant = "contained", isLoading, ...restProps }: ButtonProps) {
     return (
          <TouchableOpacity
               disabled={isLoading}
               onPress={onPress}
               style={[defaultShortcuts(restProps),
                    styles[variant].button,
                    isLoading && disabled.button
               ]}
          >
               {isLoading ? 
                    <ActivityIndicator animating size={22} /> : <Text style={styles[variant].text}>{ children}</Text>    
          }
          </TouchableOpacity>
     )
}

const styles = {
  contained: StyleSheet.create({
    button: {
      padding: 14,
      borderRadius: 50,
      backgroundColor: "black",
    },
    text: {
      textAlign: "center",
      color: "white",
      fontSize: 18,
    },
  }),
  outlined: StyleSheet.create({
    button: {
      padding: 14,
      borderRadius: 50,
      borderColor: "darkgray",
      borderWidth: 1,
    },
    text: {
      textAlign: "center",
      color: "black",
      fontSize: 18,
    },
  }),
  ghost: StyleSheet.create({
    button: {
      padding: 14,
      borderRadius: 50,
      backgroundColor: "transparent",
    },
    text: {
      textAlign: "center",
      color: "black",
      fontSize: 18,
    },
  }),
};

const disabled = StyleSheet.create({
     button: {
          opacity: 0.5
     }
})