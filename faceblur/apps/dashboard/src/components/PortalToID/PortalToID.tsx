import {PropsWithChildren, useState, useEffect} from "react";
import ReactDOM from "react-dom";

type PortalToIDProps = PropsWithChildren<{
	id: string;
}>;

export function PortalToID({id, children}: PortalToIDProps) {
	const [targetElement, setTargetElement] = useState<HTMLElement | null>(null);

	// Ensure the target element exists or create it dynamically
	useEffect(() => {
		const element = document.getElementById(id);
		if (element) {
			setTargetElement(element);
		}
	}, [id]);

	if (!targetElement) {
		return null; // Wait until the target element is available
	}

	return ReactDOM.createPortal(
		children,
		targetElement
	)
}
