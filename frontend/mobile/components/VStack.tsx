import { Stack, StackProps } from "./Stack";

interface VStackProps extends StackProps{}

export function VStack(props: VStackProps) {
     return (
          <Stack {...props} direction="column" />
     )
}