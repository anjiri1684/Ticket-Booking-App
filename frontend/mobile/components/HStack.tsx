import { Stack, StackProps } from "./Stack";

interface HStackProps extends StackProps{}

export function HStack(props: HStackProps) {
     return (
          <Stack {...props} direction="row" />
     )
}