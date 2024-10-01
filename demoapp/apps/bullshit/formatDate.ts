export const formatDate = (date: string | null) => {
    if (!date) {
        return '';
    }

    const dateInstance = new Date(date);

    const month = (dateInstance.getMonth() + 1).toString().padStart(2, '0');
    const day = dateInstance
        .getDate()
        .toString()
        .padStart(2, '0');
    const year = dateInstance.getFullYear();

    return `${day}.${month}.${year}`;
};
