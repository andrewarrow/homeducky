div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div mt-6 text-2xl text-center
        div bg-purple-500 rounded-lg p-3
          HomeDucky.com - One Product a Day
        form id=add
          div mt-9
            input input input-primary id=asin placeholder=asin autofocus=true
          div mt-3
            input input input-primary id=title placeholder=title autofocus=true
          div mt-6
            input id=add-button type=submit btn btn-primary value=Add
        div
          {{ range $i, $item := .items }}
            div 
              div mt-3 flex space-x-6
                div w-9
                  {{ add $i 1 }}.
                div
                  <img src="{{$item.photo}}" class="w-32 h-32 rounded-lg object-cover"/>
                div
                  p
                    {{ $item.asin }}
                  p text-sm text-gray-600 w-64
                    {{ $item.original_title }}
                  p text-sm text-red-600
                    {{ $item.scheduled_for_human }}
          {{ end }}
