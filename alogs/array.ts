const bubbleSort = <T extends number>(arr: T[]): T[] => {
	for (let i = 0; i < arr.length; i++) {
		for (let j = i; j < arr.length; j++) {
			if (arr[i] > arr[j]) {
				const tmp = arr[i];
				arr[i] = arr[j];
				arr[j] = tmp;
			}
		}
	}

	return arr;
};

const merge = <T extends number>(arr1: T[], arr2: T[]): T[] => {
	let result: T[] = [];

	while(arr1.length && arr2.length) {
		const element = arr1[0] < arr2[0] ? arr1.shift() : arr2.shift();
		if (element) {
			result.push(element);
		}
	}

	return [...result, ...arr1, ...arr2];
};

const mergeSort = <T extends number>(arr: T[]): T[] => {
	if (arr.length <= 1) {
		return arr;
	}

	const middleIndex = arr.length / 2;

	let leftPart = arr.splice(0, middleIndex);
	let rightPart = arr;

	return merge(mergeSort(leftPart), mergeSort(rightPart));
};

const quickSort = <T extends number>(array: T[], indexStart: number, indexEnd: number): T[] => {
	if (array.length > 1 && indexStart < indexEnd) {
		const pivotIndex = partition(array, indexStart, indexEnd);

		// in the left part there is no elements that are greater than array[pivotIndex]
		// that's why repeat the operation for that part
		if (indexStart < pivotIndex - 1) {
			quickSort(array, indexStart, pivotIndex - 1);
		}

		// in the right part there is no elements that are smaller than array[pivotIndex]
		// that's why repeat the operation for that part
		if (pivotIndex < indexEnd) {
			quickSort(array, pivotIndex, indexEnd);
		}
	}

	return array;
};

const swap = <T extends number>(array: T[], indexOne: number, indexTwo: number) => {
	let tmp = array[indexOne];
	array[indexOne] = array[indexTwo];
	array[indexTwo] = tmp;
};

const partition = <T extends number>(array: T[], indexStart: number, indexEnd: number): number => {
	// get an element in between
	const chosenElement = array[Math.floor((indexStart + indexEnd) / 2)];

	// take indexes of start and end
	let left = indexStart;
	let right = indexEnd;

	// start moving indexes toward each other until they meet
	while(left <= right) {
		// moving to the left and search for the first element which is smaller then chosenElement
		while(array[left] < chosenElement) {
			left += 1;
		}

		// moving to the right and search for the first element which is greater then chosenElement
		while(array[right] > chosenElement) {
			right -= 1;
		}

		// if indexes didn't go one over another
		if (left <= right) {
			// then swap the elements
			swap(array, left, right);
			// move indexes forward
			left += 1;
			right -= 1;
		}
	}

	// return an index on which the cycle has ended
	return left;
};


const binarySearchIterative = <T extends number>(input: T[], element: T): number => {
	let startIndex = 0;
	let endIndex = input.length - 1;

	while(endIndex >= startIndex) {
		const midIndex = Math.floor((endIndex + startIndex) / 2);
		const midElement = input[midIndex];

		if (midElement === element) {
			return midIndex;
		}

		if (midElement > element) {
			endIndex = midIndex - 1;
		}

		if (midElement < element) {
			startIndex = midIndex + 1;
		}
	}

	return -1;
};

function main() {
	const arr = [4,1,7,5,6];

	const sortedArr = quickSort(arr, 0, arr.length - 1);
	console.log(sortedArr);
	
	console.log(binarySearchIterative(sortedArr, 5));
}

main();
