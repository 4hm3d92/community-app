<script>
    import { Table, Pagination, PaginationItem, PaginationLink, Label, Input, Row, Col, Button, Modal, Badge, Form, FormGroup } from 'sveltestrap';
    import { onMount } from "svelte";
	import { writable } from "svelte/store";
	import { createTable, Subscribe, Render, createRender } from "svelte-headless-table";
	import { addPagination, addTableFilter, addResizedColumns } from "svelte-headless-table/plugins";
	
    
    export let columns;
    export let data;
    console.log(data)
    data = writable(data);
    console.log(columns)

    //export let date = new Date();

    //let data = writable([]);

const table = createTable(data, {
    page: addPagination(),
    tableFilter: addTableFilter(),
});
columns = table.createColumns(columns.map((elem) => 
           table.column({ 
                header: elem.replaceAll("_"," ").toUpperCase(),
                accessor: elem,
                cell: ({ value }) => {
                    if (typeof value === "boolean") {
                        return createRender(Input, {
                            type: "checkbox",
                            checked: value,
                            onclick: "return false"
                            //name: `${value.firstName} ${value.lastName}`,
                    })
                    }
                    else {
                        return value
                    }
                }

            })
         ))
/* columns = table.createColumns([
    table.column({
        header: "Mem ID",
        accessor: "id",
    }),
    table.column({
        header: "First Name",
        accessor: "first_name",
    }),
    table.column({
        header: "Middle Name",
        accessor: "middle_name",
    }),
    table.column({
        header: "Last Name",
        accessor: "last_name",
    }),
    table.column({
        header: "Gender",
        accessor: "gender",
    }),
    table.column({
        header: "Date of Birth",
        accessor: "dob",
    }),
]);
*/

const { headerRows, pageRows, tableAttrs, tableBodyAttrs, pluginStates } = table.createViewModel(columns);
const { pageIndex, pageCount, pageSize, hasNextPage, hasPreviousPage } = pluginStates.page;
export const { filterValue } = pluginStates.tableFilter;


</script>




<Table hover {...$tableAttrs} class="border-secondary">
    <thead class="table-dark">
        {#each $headerRows as headerRow (headerRow.id)}
            <Subscribe rowAttrs={headerRow.attrs()} let:rowAttrs>
                <tr {...rowAttrs}>
                    {#each headerRow.cells as cell (cell.id)}
                        <Subscribe
                            attrs={cell.attrs()}
                            let:attrs
                            props={cell.props()}
                            let:props
                        >
                            <th {...attrs}>
                                <Render of={cell.render()} />
                            </th>
                        </Subscribe>
                    {/each}
                </tr>
            </Subscribe>
        {/each}
    </thead>
    <tbody {...$tableBodyAttrs}>
        {#each $pageRows as row (row.id)}
            <Subscribe rowAttrs={row.attrs()} let:rowAttrs>
                <tr {...rowAttrs}>
                    {#each row.cells as cell (cell.id)}
                        <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
                            <td {...attrs} class:matches={props.tableFilter.matches}>
                                <Render of={cell.render()} />
                                
                            </td>
                        </Subscribe>
                    {/each}
                </tr>
            </Subscribe>
        {/each}
    </tbody>
</Table>

<div style="display: flex; justify-content: center">
	<Pagination ariaLabel="Page navigation">
		<PaginationItem disabled={!$hasPreviousPage} >
		  <PaginationLink  href="#" on:click={() => $pageIndex--}>
			Previous
		  </PaginationLink>
			
		</PaginationItem>
		<PaginationItem disabled >
			<PaginationLink >
				{$pageIndex + 1} out of {$pageCount}
			</PaginationLink>
		</PaginationItem>
		<PaginationItem disabled={!$hasNextPage} >
		  <PaginationLink next href="#" on:click={() => $pageIndex++}>
		    Next	
		</PaginationLink>
		</PaginationItem>
	</Pagination>
    <br>
</div>

<Label for="page-size">Page size</Label>
<Input id="page-size" type="number" min={1} bind:value={$pageSize} />



<style>
    .matches {
      background: rgba(46, 196, 182, 0.8);
    }
  
  </style>