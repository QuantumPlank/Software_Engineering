from logger import ConsoleLogger, FilenameDecorator, LevelDecorator, TimestampDecorator


def main() -> None:
    basic = ConsoleLogger()
    basic.log("Basic Logger")

    decorated = TimestampDecorator(
        logger=FilenameDecorator(
            logger=LevelDecorator(
                logger=basic,
                level="Debug",
            ),
            filename="dummy.txt",
        ),
    )
    decorated.log("Decorated Logger")


if __name__ == "__main__":
    main()
